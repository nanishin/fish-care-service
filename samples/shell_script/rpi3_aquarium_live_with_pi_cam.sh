#!/bin/bash

SERVER_URL=rtmp://YOUR_RTMP_SERVER_URL
STREAM_KEY=YOUR_RTMP_STREAM_KEY
BASE_BGM_AUDIO=YOUR_M4A_FILE_PATH
AUDIO_STREAM_PIPE=/tmp/audio_stream.m4a

rm ${AUDIO_STREAM_PIPE}
mkfifo ${AUDIO_STREAM_PIPE}

# 30FPS, 5Mbps, 1080p Video, 128K Audio
/usr/local/bin/ffmpeg -y -stream_loop -1 -i ${BASE_BGM_AUDIO} -c:a copy -movflags +faststart -f nut /tmp/audio_stream.m4a | /usr/bin/raspivid -o - -t 0 -vf -hf -fps 30 -b 5000000 -w 1920 -h 1080 -awb sun | /usr/local/bin/ffmpeg -thread_queue_size 512 -i /tmp/audio_stream.m4a -f h264 -thread_queue_size 512 -i - -vcodec copy -acodec aac -ab 128k -g 50 -strict experimental -f flv ${SERVER_URL}/${STREAM_KEY}
