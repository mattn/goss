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

    goss put file.txt sakura://<bucket>/path/to/
    goss put file.txt sakura://<bucket>/path/to/another-name.txt

Download

    goss get sakura://<bucket>/path/to/file.txt

Cat

    goss cat sakura://<bucket>/path/to/file.txt
