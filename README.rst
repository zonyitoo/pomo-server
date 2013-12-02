Pomo Server
-----------

A pomo server for Pomo App.

Dependences
===========

* Revel_

.. _Revel: http://robfig.github.io/revel

Usage
=====

.. code:: bash

	cd pomo-server
	export GOPATH=$PWD  # set current directory as $GOPATH
	go get github.com/robfig/revel/revel  # Get Revel web framework
	export PATH=$PWD/bin:$PATH  # Add bin directory to $PATH
	revel run pomo-server