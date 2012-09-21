#!/bin/bash

PUML="java -jar `realpath ~/apps/plantuml/plantuml.jar` -o `pwd`"
#PUML=$PUML" -tsvg"

$PUML *.diagram
