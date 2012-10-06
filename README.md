# Braille Printer

Source of [Braille Printer](http://braille-printer.appspot.com/)

## Dependency

It uses [hangul](https://github.com/suapapa/go_hangul)
and [braille](go_braille) as submodule.
Issue following command to update them.

    $ git submodule init
    $ git submodule update

## Test

After download and extract Google App Engine SDK for Go
from [Download](https://developers.google.com/appengine/downloads),
run test server in localhost;

    $ path_to_appengine_sdk_for_go/dev_appserver.py ./

And, connect to `localhost:8080` with browser


## APIs

    Host: http://braille-printer.appspot.com/

    POST /braille
      input: 변환할스트링
      lang: auto|ko|en

    POST /printq/add
      input: 변환할스트링
      lang: auto|ko|en
      key: examplekey

    GET  /printq/list?type=label|paper|all&key=examplekey
      { 
        [
          {"qid": 1234, 
           "type": "label|paper"},
          {"qid": 5678, 
           "type": "label|paper"}
        ]
      }

    GET  /printq/item?qid=1234&format=text|svg&key=examplekey
      {
        "origin": "변환할 문자열",
        "result": "변환된 점자"
      }

    POST /printq/update?qid=1234&status=1&key=examplekey
