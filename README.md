# Lonker
## Description
A FOSS (free and open source) Link shortening service that isn't online yet.

## Api
There is many examples like the said file in `/assets/scripts/api.js` that contains functions for interacting with the lonker api. <br>
If those don't satisfy you enough here is some other examples in different languages.<br>
Python3
```py
import requests
r = requests.post('http://104.243.44.32:8080/api/v0/upload', json={"url": "https://youtube.com/"})
# if it was succesfully
if r.status_code == 200 {
    shortenedlink = r.json()["data"]["url"]
}
```
Php (yes I did do that)
```php
$json = ["url" => "https://youtube.com/"];
$ch = curl_init("http://104.243.44.32:8080/api/v0/upload");
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($json));
curl_setopt($ch, CURLOPT_HTTPHEADER, ["Content-type: application/json"];
$data = curl_exec($ch);
curl_close($ch);
echo $data;
```
C#
```cs
var url = "http://104.243.44.32:8080/api/v0/upload";

var httpRequest = (HttpWebRequest)WebRequest.Create(url);
httpRequest.Method = "POST";

httpRequest.ContentType = "application/json";

var data = "{\"url\": \"https://youtube.com/\"}";

using (var streamWriter = new StreamWriter(httpRequest.GetRequestStream()))
{
   streamWriter.Write(data);
}

var httpResponse = (HttpWebResponse)httpRequest.GetResponse();
using (var streamReader = new StreamReader(httpResponse.GetResponseStream()))
{
   var result = streamReader.ReadToEnd();
}

Console.WriteLine(httpResponse.StatusCode);
```
Java
```java
URL url = new URL("http://104.243.44.32:8080/api/v0/upload");
HttpURLConnection http = (HttpURLConnection)url.openConnection();
http.setRequestMethod("POST");
http.setDoOutput(true);
http.setRequestProperty("Content-Type", "application/json");

String data = "{\"url\": \"https://youtube.com/\"}";

byte[] out = data.getBytes(StandardCharsets.UTF_8);

OutputStream stream = http.getOutputStream();
stream.write(out);

System.out.println(http.getResponseCode() + " " + http.getResponseMessage());
http.disconnect();
```
Bash
```bash
#!/bin/bash
curl -X POST http://104.243.44.32:8080/api/v0/upload -H "Content-Type: application/json" -d "{\"url\": \"https://youtube.com/\"}"
```
