# Changelog

## v0.3.0

- Added dependency on [fireworks](https://github.com/JonnyOrman/fireworks) `v0.1.0`.
- Removed application building, configuration and request handling utilities that are now provided by `fireworks`.
- Removed the following types and functions:
    - Application
    - ApplicationConfigurationLoader
    - Configuration
    - ConfigurationFilePathProvider
    - ConfigurationFileReader
    - ConfigurationJsonFileReader
    - ConfigurationLoader
    - ConfigurationValueProvider
    - CreateConfigurationValueProvider
    - DataDeserialiser
    - DataReader
    - EnvironmentValueProvider
    - FilePathProvider
    - FileReader
    - GinPubSubBodyReader
    - GinRouterBuilder
    - HttpRequestBodyDataReader
    - IoReaderGenerator
    - IoutilReader
    - JsonDataDeserialiser
    - JsonReader
    - JsonValueProvider
    - PubSubBody
    - PubSubBodyReader
    - Reader
    - RequestHandler
    - RouterBuilder
    - ValueProvider

## v0.2.0

- Configuration can now be loaded from the environment instead of a config file.

## v0.1.2

- Internal refactor to split code up into smaller units.

## v0.1.1

- Internal refactor to improve code reuse.

## v0.1.0

- Receive Pub/Sub push messages and insert their data into a Firestore collection.