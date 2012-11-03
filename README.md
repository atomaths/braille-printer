# Braille Printer

Source of [Braille Printer](http://braille-printer.appspot.com/)

## Clients

- [CLI client](https://github.com/dalinaum/braille-printer-client)
- [Chrome Extension](https://github.com/golanger/braille-printer-chrome-extension)

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

- Host: http://braille-printer.appspot.com/

### 사용자 

점자 변환

    POST /braille
      input: 변환할스트링
      lang: auto|ko|en
      format: svg|text

점자 변환 및 클라우드 프린팅 요청

    POST /printq/add
      input: 변환할스트링
      lang: auto|ko|en
      key: examplekey_user

### 프린터

인쇄 대기중인 목록 가져오기

    GET  /printq/list?type=label|paper|all&key=examplekey_printer
      { 
        [
          {"qid": 1234, 
           "type": "label|paper"},
          {"qid": 5678, 
           "type": "label|paper"}
        ]
      }

특정 인쇄 아이디의 내용 가져오기 

    GET  /printq/item?qid=1234&format=text|svg&key=examplekey_printer
      {
        "origin": "변환할 문자열",
        "result": "변환된 점자"
      }

특정 인쇄 아이디의 상태(인쇄됨) 변경

    POST /printq/update?qid=1234&status=1&key=examplekey_printer

