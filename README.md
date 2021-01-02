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
## Contributing
At the moment I am looking for frontend developers for this project, so any help on the html design end is much much appreciated. <br>
For inquires contact me on discord Jviguy#2975
