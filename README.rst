Pomo Server
-----------

A pomo server for Pomo App.

Dependences
===========

* Revel_

.. _Revel: http://robfig.github.io/revel

* mgo_

.. _mgo: http://labix.org/mgo

* gomemcache_

.. _gomemcache: https://github.com/bradfitz/gomemcache

Usage
=====

.. code:: bash

	cd pomo-server
	export GOPATH=$PWD  # set current directory as $GOPATH
	go get github.com/robfig/revel/revel  # Get Revel web framework
	export PATH=$PWD/bin:$PATH  # Add bin directory to $PATH
	go get labix.org/v2/mgo # Get mgo
	go get github.com/bradfitz/gomemcache # Get gomemcache
	revel run pomo-server