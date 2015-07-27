#!/usr/bin/env python3

import ctypes

testlib = ctypes.CDLL('./target/release/libembed.dylib')

testlib.process()

print("done")
