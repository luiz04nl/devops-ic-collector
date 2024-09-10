# FROM maven:3.9.7-eclipse-temurin-8
# FROM maven:3.9.7-eclipse-temurin-11
FROM maven:3.9.7-eclipse-temurin-22

WORKDIR /app

RUN apt-get update
RUN apt-get install -y git
RUN apt-get install -y unzip

# RUN curl -L -o ck-0.7.0.zip https://github.com/mauricioaniche/ck/archive/refs/tags/ck-0.7.0.zip
# RUN unzip ck-0.7.0.zip -d ck
# WORKDIR /app/ck/ck-ck-0.7.0
# # RUN mvn compile
# RUN mvn clean compile package

RUN curl -L -o /app/ck-0.7.0-jar-with-dependencies.jar https://repo1.maven.org/maven2/com/github/mauricioaniche/ck/0.7.0/ck-0.7.0-jar-with-dependencies.jar

COPY ./docker/startup.sh /usr/local/bin/startup.sh
RUN chmod +x /usr/local/bin/startup.sh

ENTRYPOINT ["/usr/local/bin/startup.sh"]
