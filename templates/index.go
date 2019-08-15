package templates

const Index string = `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Goupfile - Secure and anonymous file upload from the command line</title>
	<link rel="stylesheet" href="https://unpkg.com/@bafs/mu@0.3/mu.min.css" />
	<style>
		body {
			font-family: -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen-Sans,Ubuntu,Cantarell,"Helvetica Neue",sans-serif;
		}
		small {
			font-size: 0.7rem !important;
		}
	</style>
</head>
<body>
	<h1>Goupfile <small>alpha</small></h1>
	<h2>What is it?</h2>
	<p>Goupfile is secure and anonymous file upload from the command line. It allows you to quickly share files and get a links for them without leaving your terminal.</p>
	<h2>How does it work?</h2>
	<p>As of now, Goupfile will be a service that you interact with through the goup CLI tool.</p>
	<p>For example, you might do this:</p>
	<pre>
$ goup file.txt
Uploading file...

	https://goupfile.com/aw9kzm

$ goup download aw9kzm
</pre>
	<p>However, while the CLI tool is being developed, you can send an HTTP POST request with a multipart form body to upload a file. The key should be named 'file'.</p>
	<hr>
	<a href="https://github.com/goupfile">GitHub</a>
</body>
</html>
`
