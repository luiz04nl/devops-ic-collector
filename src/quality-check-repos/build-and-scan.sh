#!/usr/bin/env bash

REPOSITORY=$1
WITH_BUILD=${2:-true}
$SCRIPT_BASE_PATH=$3

echo "WITH_BUILD: $WITH_BUILD"

if [ "$WITH_BUILD" = "true" ];
then
  echo "WITH_BUILD: Using Full Build"
else
  echo "WITH_BUILD: Sipped Full Build"
fi

cd $SCRIPT_BASE_PATH
cd ../../repos/
cd $REPOSITORY

BUILD_PATH=`pwd`
echo "BUILD_PATH: $BUILD_PATH"

mkdir -p $SCRIPT_BASE_PATH/../../out/quality-check-repos

SONAR_URL="http://localhost:9000"
SONAR_USERNAME="admin"
SONAR_PASSWORD="sonar"
SONAR_CREDENTIALS="${SONAR_USERNAME}:${SONAR_PASSWORD}"
ENCODED_CREDENTIALS=$(echo -n "$SONAR_CREDENTIALS" | base64)

cat <<EOF > sonar-project.properties
sonar.projectKey=$REPOSITORY
sonar.projectName=$REPOSITORY
sonar.projectVersion=1.0
sonar.sourceEncoding=UTF-8
sonar.host.url=$SONAR_URL
sonar.login=$SONAR_USERNAME
sonar.password=$SONAR_PASSWORD
sonar.verbose=true
EOF

projectTypeVersion=""
projectType=""
analysisSuccess=undefined
filesAtRootDir=$(ls .)

mkdir -p target/classes
mkdir -p target/test-classes


if [ -f "pom.xml" ]; then
  cat <<EOF >> sonar-project.properties
sonar.language=java
sonar.java.binaries=target/classes,target/test-classes
EOF

  projectType="maven"
  distributionUrl=`cat .mvn/wrapper/maven-wrapper.properties | grep distributionUrl`
  splitA=(${distributionUrl//"distributionUrl=https://repo.maven.apache.org/maven2/org/apache/maven/apache-maven/"/ })
  splitB=(${splitA//"-all.zip"/ })
  splitC=(${splitB//"-bin.zip"/ })

  projectTypeVersion=(${splitC//"/apache-maven-"/ })

  defaultProjectTypeVersion="3.9.6-eclipse-temurin-11"

  if [ -z "$projectTypeVersion" ]; then
      projectTypeVersion=$defaultProjectTypeVersion
  fi

  image_exists() {
    docker image inspect "$1" > /dev/null 2>&1
  }

  if ! image_exists "$projectTypeVersion"
  then
    projectTypeVersion=$defaultProjectTypeVersion
  fi

  if [ "$WITH_BUILD" = "true" ]
  then
    MSYS_NO_PATHCONV=1 docker run --rm -v "$(pwd)":/usr/src/mymaven -w /usr/src/mymaven maven:$projectTypeVersion /bin/bash -c "git config --global --add safe.directory /usr/src/mymaven && mvn compile" 2>&1
  fi

elif [ -f "build.gradle" ]; then
  mkdir -p build/classes/java/main
  mkdir -p build/classes/java/test

  cat <<EOF >> sonar-project.properties
sonar.language=java
sonar.java.binaries=build/classes/java/main,build/classes/java/test
EOF

  projectType="gradle"
  if [ -f "app/src/main/AndroidManifest.xml" ]; then
    projectType="gradle-android"
  fi

  distributionUrl=`cat gradle/wrapper/gradle-wrapper.properties  | grep distributionUrl`
  splitA=(${distributionUrl//"distributionUrl=https\://services.gradle.org/distributions/gradle-"/ })
  splitB=(${splitA//"-all.zip"/ })
  projectTypeVersion=(${splitB//"-bin.zip"/ })

  defaultProjectTypeVersion="8.7.0-jdk8"

  if [ -z "$projectTypeVersion" ]; then
      projectTypeVersion=$defaultProjectTypeVersion
  fi
  image_exists() {
    docker image inspect "$1" > /dev/null 2>&1
  }
  if ! image_exists "$projectTypeVersion"
  then
    projectTypeVersion=$defaultProjectTypeVersion
  fi

  if [ "$WITH_BUILD" = "true" ]
  then
    MSYS_NO_PATHCONV=1 docker run --rm -u gradle -v `pwd`:/home/gradle/project -w /home/gradle/project gradle:$projectTypeVersion /bin/bash -c "git config --global --add safe.directory /home/gradle/project && gradle build" 2>&1
  fi

elif  [ -f "build.xml" ]; then
  cat <<EOF >> sonar-project.properties
sonar.language=java
sonar.java.binaries=**/**/classes
EOF

  projectType="ant"
  distributionUrl=`cat build.xml | grep antversion | xargs`
  splitA=(${distributionUrl/"<antversion atleast=\""/ })
  projectTypeVersion=(${splitA//"\""/ })

  defaultProjectTypeVersion="latest"

  if [ -z "$projectTypeVersion" ]; then
      projectTypeVersion=$defaultProjectTypeVersion
  fi
  image_exists() {
    docker image inspect "$1" > /dev/null 2>&1
  }
  if ! image_exists "$projectTypeVersion"
  then
    projectTypeVersion=$defaultProjectTypeVersion
  fi

  if [ "$WITH_BUILD" = "true" ]
  then
    MSYS_NO_PATHCONV=1 docker run --rm -v "$(pwd)":/workspace -w /workspace bitnami/ant:$projectTypeVersion /bin/bash -c "git config --global --add safe.directory/workspace && ant build" 2>&1
  fi

elif  [ -f "package.json" ]; then
  projectType="JavaScript"
  if [ -f "tsconfig.json" ]; then
    projectTypeVersion="TypeScript"

    cat <<EOF >> sonar-project.properties
sonar.sources=src
sonar.language=js,ts
sonar.inclusions=**/*.js,**/*.ts
sonar.exclusions=node_modules/**,dist/**
sonar.typescript.tsconfigPath=tsconfig.json
sonar.javascript.lcov.reportPaths=coverage/lcov.info
sonar.typescript.lcov.reportPaths=coverage/lcov.info
EOF
  else
      projectTypeVersion="JavaScript"
    cat <<EOF >> sonar-project.properties
sonar.sources=src
sonar.language=js
sonar.inclusions=**/*.js
sonar.exclusions=node_modules/**,dist/**
sonar.javascript.lcov.reportPaths=coverage/lcov.info
EOF
  fi

elif [ -f "requirements.txt" ] || [ -f "setup.py" ] ||  -f "pyproject.toml" ]; then
  projectType="Python"
  projectTypeVersion="Python"
  cat <<EOF >> sonar-project.properties
sonar.sources=.
sonar.language=py
sonar.python.version=3.x
sonar.python.coverage.reportPaths=coverage.xml
sonar.exclusions=**/tests/**
EOF

else
  projectType="other"
  projectTypeVersion=""

  if find . -name "*.java" -not -name "test.java" -print -quit | grep -q '.'; then

    cat <<EOF >> sonar-project.properties
sonar.language=java
sonar.java.binaries=target/classes/
EOF

    projectType="undefined"
    projectTypeVersion="undefined"

    cat <<EOF > custom-pom.xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.devops</groupId>
    <artifactId>test</artifactId>
    <version>1.0-SNAPSHOT</version>

    <properties>
        <maven.compiler.source>1.8</maven.compiler.source>
        <maven.compiler.target>1.8</maven.compiler.target>
    </properties>

    <dependencies>
    </dependencies>

    <build>
        <plugins>
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-compiler-plugin</artifactId>
                <version>3.8.1</version>
                <configuration>
                    <source>\${maven.compiler.source}</source>
                    <target>\${maven.compiler.target}</target>
                </configuration>
            </plugin>
            <plugin>
                <groupId>org.sonarsource.scanner.maven</groupId>
                <artifactId>sonar-maven-plugin</artifactId>
                <version>3.7.0.1746</version>
            </plugin>
        </plugins>
    </build>
</project>
EOF

    srcDir="src/main/java"
    if [ -d "$srcDir" ]; then
        rm -r $srcDir
    fi
    mkdir -p $srcDir

    find . -name "*.java" -exec cp {} $srcDir/ \;

  mavenVersion="3.8.1"
  if [ "$WITH_BUILD" = "true" ]
  then
    MSYS_NO_PATHCONV=1 docker run --rm -v "$(pwd)":/usr/src/mymaven -w /usr/src/mymaven maven:$mavenVersion /bin/bash -c "git config --global --add safe.directory /usr/src/mymaven && mvn -f custom-pom.xml compile" 2>&1
  fi

  else
    projectType="other"
    projectTypeVersion=""
  fi
fi

analysisSuccess="undefined"
if grep -q "EXECUTION SUCCESS" $SCRIPT_BASE_PATH/../../out/quality-check-repos/$REPOSITORY.out.txt; then
    analysisSuccess=true
else
    analysisSuccess=false
fi

export projectType=$projectType
export projectTypeVersion=$projectTypeVersion
export analysisSuccess=$analysisSuccess
export filesAtRootDir=$filesAtRootDir

export SONAR_USERNAME=$SONAR_USERNAME
export SONAR_PASSWORD=$SONAR_PASSWORD
export SONAR_URL=$SONAR_URL
export SONAR_CREDENTIALS=$SONAR_CREDENTIALS
export ENCODED_CREDENTIALS=$ENCODED_CREDENTIALS

if [ "$WITH_BUILD" = "true" ]
then
  echo "SONAR_URL: $SONAR_URL"
  echo "REPOSITORY: $REPOSITORY"

  sonar-scanner --version
  sonar-scanner -X

  # #MSYS_NO_PATHCONV=1 docker run --rm -e SONAR_HOST_URL="$SONAR_URL" -v "$(pwd):/usr/src" sonarsource/sonar-scanner-cli --version 2>&1
  # #MSYS_NO_PATHCONV=1 docker run --rm -e SONAR_HOST_URL="$SONAR_URL" -v "$(pwd):/usr/src" sonarsource/sonar-scanner-cli -X 2>&1
fi