#!/bin/bash

function validate_single_arg () {
  local val=$1
  local arg=${l_arg:1}
  local res=${arg//[^$val]}
  local cant=${#res}
  validate_single_arg_item_status=$2
  validate_single_arg_error=false

  if [ $cant -eq 1 ]
  then 
    if [ $validate_single_arg_item_status = false ]
    then 
      validate_single_arg_item_status=true
    else
      validate_single_arg_error=true
    fi
  elif [ $cant -gt 1 ]
  then
    validate_single_arg_error=true
  fi
}

function validate_args () {
  local l_args=$1

  local validate_single_arg_item_status
  local validate_single_arg_error

  local l_arg_error=false
  local l_comb_error=false
  local l_arg
  local l_init_dev_env=false
  local l_start_server=false
  local l_start_prompt=false
  local l_down_istance=false
  local l_show_help=false

  if [[ "$l_args" == "" ]]
  then
    l_arg_error=true
  else
    for l_arg in ${args[@]}
    do
      if [[ ${l_arg:0:1} = '-' ]]
      then
        if [[ ${l_arg:1} =~ [^i|s|p|d|h] ]] || [[ ${l_arg:1} = '' ]]
        then 
          l_arg_error=true
        else
          if [ $l_arg_error = false ]
          then
            # se valida el parametro i 
            validate_single_arg 'i' $l_init_dev_env
            l_init_dev_env=$validate_single_arg_item_status
            l_arg_error=$validate_single_arg_error
          fi
          if [ $l_arg_error = false ]
          then
            # se valida el parametro s 
            validate_single_arg 's' $l_start_server
            l_start_server=$validate_single_arg_item_status
            l_arg_error=$validate_single_arg_error
          fi
          if [ $l_arg_error = false ]
          then
            # se valida el parametro p 
            validate_single_arg 'p' $l_start_prompt
            l_start_prompt=$validate_single_arg_item_status
            l_arg_error=$validate_single_arg_error
          fi
          if [ $l_arg_error = false ]
          then
            # se valida el parametro d 
            validate_single_arg 'd' $l_down_istance
            l_down_istance=$validate_single_arg_item_status
            l_arg_error=$validate_single_arg_error
          fi
          if [ $l_arg_error = false ]
          then
            # se valida el parametro h 
            validate_single_arg 'h' $l_show_help
            l_show_help=$validate_single_arg_item_status
            l_arg_error=$validate_single_arg_error
          fi
        fi
      else
        l_arg_error=true
        echo "paso por aca"
      fi
      if [ $l_arg_error = true ]; then break; fi
    done
    
    if [ $l_arg_error = true ]
    then 
      arg=$l_arg
    else 
      arg=''
      if [ $l_show_help = true ]
      then
        if [ $l_init_dev_env = true ]; then l_comb_error=true; fi
        if [ $l_start_server = true ]; then l_comb_error=true; fi
        if [ $l_start_prompt = true ]; then l_comb_error=true; fi
        if [ $l_down_istance = true ]; then l_comb_error=true; fi
      fi
      if [ $l_init_dev_env = true ]
      then
        if [ $l_start_server = true ]; then l_comb_error=true; fi
        if [ $l_start_prompt = true ]; then l_comb_error=true; fi
        if [ $l_down_istance = true ]; then l_comb_error=true; fi
      fi
      if [ $l_start_server = true ]
      then
        if [ $l_start_prompt = true ]; then l_comb_error=true; fi
      fi
    fi
  fi
  
  args_error=$l_arg_error
  comb_error=$l_comb_error

  init_dev_env=$l_init_dev_env
  start_server=$l_start_server
  start_prompt=$l_start_prompt
  down_istance=$l_down_istance
  show_help=$l_show_help
}

function action_down_istance () {
  docker-compose down
  if [ $(docker images -f DANGling=true -q) ]
  then
    docker rmi $(docker images -f DANGling=true -q)
  fi
}

function action_init_dev_env () {
  if ! [ "$(ls -laF | grep react/)" ]
  then
    docker-compose up --build react-init
    action_down_istance
  else
    echo "Environment is already initialized!"
  fi
  if ! [ "$(ls -laF | grep .git/)" ]
  then
    git init
  fi  
}

function action_start_server () {
  docker-compose up --build react-dev-server
}

function action_start_prompt () {
  if  ! ( docker-compose exec react-dev bash )
  then 
    docker-compose up --build -d react-dev
    docker-compose exec react-dev bash
  fi
}

function action_show_help () {
  echo ""
  echo " Options:"
  echo ""
  echo "  -i   Run docker-compose to initialize dev environmenet (can't be combined with other options)."
  echo "  -s   Run docker-compose to start server."
  echo "  -p   Run docker-compose to start bash prompt."
  echo "  -d   Close instance, can be conbined to be executed after use a run option."
  echo "  -h   Show help (can not be combined with other options)."
  echo ""
}

args=( $@ )

validate_args $args

if [ $args_error = true ]
then
  echo "invalid option(s): '$@', use -h option for help."
elif [ $comb_error = true ]
then
  echo "invalid options combination, use -h option for help."
else
  if [ $init_dev_env = true ]; then action_init_dev_env; fi
  if [ $start_server = true ]; then action_start_server; fi
  if [ $start_prompt = true ]; then action_start_prompt; fi
  if [ $down_istance = true ]; then action_down_istance; fi
  if [ $show_help = true ]; then action_show_help; fi
fi