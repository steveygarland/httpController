#!/usr/bin/env python
import socket

from http_parser.http import HttpStream
from http_parser.reader import SocketReader

from http_parser.util import b

def main():
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.bind((socket.gethostname(), 8080))
    s.listen(5)
    try:
        while True:
            p = HttpStream(SocketReader(s))
    finally:
        s.close()

if __name__ == "__main__":
    main()

print p
