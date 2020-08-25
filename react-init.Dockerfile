FROM debian:oldstable-20200514-slim

#Installing sudo
RUN apt update && apt install -y sudo curl

#Installing nodejs
RUN curl -sL https://deb.nodesource.com/setup_14.x | bash -
RUN apt update && apt install -y nodejs

#Creating user developer
RUN useradd -ms /bin/bash developer --uid 1000
RUN echo "developer ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

#Changing to user developer
USER developer
WORKDIR /home/developer/app
RUN sudo chown -R developer:developer /home/developer

#Installing nodejs
RUN curl -o- -L https://yarnpkg.com/install.sh | bash
ENV PATH /home/developer/.yarn/bin:$PATH