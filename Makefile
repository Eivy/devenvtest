PHONY: image deploy

image: image
	gcloud builds submit --tag gcr.io/aptitude-bulb-308204/api

deploy: 
	gcloud run deploy --image gcr.io/aptitude-bulb-308204/api --platform managed
