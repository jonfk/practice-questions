#!/usr/bin/env python3

import threading



def run():
    count = 0
    for i in range(0,5000000):
        count += 1

    print(count)


for i in range(0,10):
    mythread = threading.Thread(name="mythread", target=run)
    mythread.start()
    mythread.join()
