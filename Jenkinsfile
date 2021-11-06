pipeline {
  agent any
  
  tools {
        go 'Go 1.17.3'
  }
  environment {
      GO114MODULE = 'on'
      CGO_ENABLED = 0 
      GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
  }
  
  stages {
    stage ('Run Unit Test') {
        steps {
          script {
            try {
              withEnv(["PATH+GO=${GOPATH}/bin"]){
                sh 'go test ./... -v'
              }
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
