# mad

A HTTP server that serves no content.

# Usage

This server requires to be activated by means of Systemd socket activation. To
test this, run: `/lib/systemd/systemd-activate -l 127.0.0.1:8076 ./bin/mad`
which will start the server on the localhost at port 8076.

# License

This code is licensed under Apache License v2. Copyright (c) 2015, Jan Willem
Janssen.


