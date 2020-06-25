Cloud speech to textを利用して、暴言を検知、その後、HTMLでその数を表示する実験装置です。

Go側
* echoでWebsocketサーバー

Python
* マイクの取得


* https://cloud.google.com/speech-to-text/docs/quickstart-client-libraries?hl=ja


```
go get 
go run main.go
```


```
#python3 -m venv /path/to/new/virtual/environment
python3 -m venv venv
```

https://docs.python.org/ja/3/library/venv.html
https://people.csail.mit.edu/hubert/pyaudio/
brew install portaudio
pip install pyaudio