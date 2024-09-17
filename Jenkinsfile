pipeline {
	agent any 

	stages {
		stage ('Lint') {
			agent {
				docker { 
					image 'golangci/golangci-lint:v1.61.0'
					reuseNode true
				}
			}
			steps {
				script {
					try {
						sh 'golangci-lint run'
						echo 'Linting Success!'
					} catch (err) {
						echo 'Lint failed'
						sh 'exit 1'
					}
				}
			}
		}
		stage ('Test') {
			agent {
				docker { 
					image 'golang:latest'
					reuseNode true
				}
			}
			steps {
				script {
					try {
						sh 'go test ./... -v'
						echo 'Test Success!'
					} catch (err) {
						echo 'Test failed'
						sh 'exit 1'
					}
				}
			}
		}
	}
}
