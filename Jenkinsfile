pipeline {
	agent any 

	stages {
		stage ('Test-Library') {
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
						sh 'go test ./... -v'
						echo 'Linting Success!'
					} catch (err) {
						echo 'Lint failed'
						sh 'exit 1'
					}
				}
			}
		}
	}
}
