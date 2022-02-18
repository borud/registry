# registry sketch

This is an updated version of the registry after some feedback from various parties. I've modeled it in terms of protocol buffers since this is a convenient way to maintain a data structure spec in a versioned manner.

The original intent was to publish the registry as a JSON file, and this is still the intent.  The grpc-gateway API definitions produce a REST+JSON API that can be used to serve the registry.

## Main changes

### Splitting sensors into devices and sensors

The main change is that we have split the registry into device and sensors.  Originally this was a flat structure that only dealt with sensors - where each sensor entry was afor a unique sensor on a unique device. This turned out to cause too much confusion, so we decided that this was not a good idea.

### Parameters

We have also added *parameters* to **devices**.  Parameters can be configuration parameters or they can represent actuation or state change commands.  For instance they can mean *"set sampling frequency of X to Y"* or they can mean *"move actuator X to position Y"*.

Parameters have numeric IDs that start at 1, so each parameter set is specific to a given **device**.  This choice was made to ensure that parameters can be encoded in as few bits as possible.
