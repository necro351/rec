rec
===

The rec package is a RESTful recommender server backed by a recommender system
that stores its reviews and users in a persistent database. It is intended as a
one-stop-shop to build a recommender system.

rec/backend
-----------

The backend package exposes a Go interface that lets you add and remove users,
add and modify reviews for users, and add and remove items for the users to
review. Items belong to categories and the same item can belong to multiple
categories.

rec/server
----------

The server package layers a RESTful API on top of the backend so that it can be
queried via HTTP.
