# FizzBuzz

FizzBuzz is a ready-to-use console application that provides tools to format output with conditions.

## Results of research

Allocation of hash map for each goroutine to avoid usage of mutex and avoid waiting for map to become free doesn't improve performance. The cost of allocation is too much, the larger the map, the more computing resources are spent.
Distribution of work between goroutines gives a significant increase in performance for multi-core processors, despite the introduction of additional mechanisms.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)