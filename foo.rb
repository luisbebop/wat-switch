puts $INPUT
req = SimpleHttp.new("https", "cloudwalk.io", 443).request("GET", "/en/press", {})
req['body']
