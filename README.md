Sakura Storage Utility
======================

Setting
-------

    SAKURA_STORAGE_USER_ID=<your service id>
    SAKURA_STORAGE_ACCESS_TOKEN=<your secret key>

How to run
----------

Listing

    goss ls sakura://<bucket>/path/to/

Upload

    goss put sakura://<bucket>/path/to/ file.txt
    goss put sakura://<bucket>/path/to/ file.txt

Download

    goss get sakura://<bucket>/path/to/file.txt

Cat

    goss cat sakura://<bucket>/path/to/file.txt
