deploy:
	gcloud functions deploy generate_pdf --entry-point GeneratePDF --runtime go120 --trigger-http --allow-unauthenticated
