FROM react:init

#Exposing ports
EXPOSE 3000

#Setting aliases for developer
RUN echo "# Custom aliases:" >> /home/developer/.bashrc  && \
    echo "alias a='clear && ls -la --color'" >> /home/developer/.bashrc  && \
    echo "alias l='ls -la --color'" >> /home/developer/.bashrc && \
    echo "alias ss='yarn start'" >> /home/developer/.bashrc

#Setting date/time to Buenos_Aires
RUN sudo cp /usr/share/zoneinfo/America/Buenos_Aires /etc/localtime

#Changing the workdir to apps dir
WORKDIR /home/developer/app/react