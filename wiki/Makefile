IMAGE_TAG=wiki-go
LOCAL_IMAGE_TAG=${IMAGE_TAG}:local-latest

build:
	GOOS=linux go build wiki.go
	docker build . -t ${LOCAL_IMAGE_TAG}

run:
	docker run -p 8080:8080 ${LOCAL_IMAGE_TAG}

push:
	docker tag ${LOCAL_IMAGE_TAG} loefesto/${IMAGE_TAG}:latest
	docker push loefesto/${IMAGE_TAG}:latest

plan:
	cd terraform &&	terraform plan

apply:
	cd terraform &&	terraform apply -auto-approve

destroy:
	cd terraform &&	terraform destroy