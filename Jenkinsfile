pipeline {
  agent any
  
  tools {
        go 'go1.17.3'
  }
  environment {
      GO111MODULE = 'on'
  }
  
  stages {
    stage ('Run Unit Test') {
        steps {
          script {
            try {
              sh 'go test ./... -v'
            } catch (e) {
              currentBuild.result = 'ABORTED'
              error("Aborting the build.")
            }
          }
        }
    }
    stage ('Build') {
        steps {
          script {
            try {
             sh 'go build'
            } catch (e) {
             currentBuild.result = 'ABORTED'
             error("Aborting the build.")
           }
          }
        }
    }
  }
}
