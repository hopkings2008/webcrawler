#!/bin/bash

nohup Xvfb :1 -ac -screen 0 1024x768x16 +extension RANDR > /dev/null 2>&1 &
