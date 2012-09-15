# Braille Printer

Source of [Braille Printer](http://braille-printer.appspot.com/)

## Dependency

It uses [hangul](https://github.com/suapapa/go_hagul)
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
