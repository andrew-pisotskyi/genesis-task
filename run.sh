#!/bin/bash
docker build -t genesis_task_a_pisotskyi .
docker run -d -it -p 8080:8080 --name=gt_a_pisotskyi genesis_task_a_pisotskyi