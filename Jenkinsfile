pipeline {
  agent any
  
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
