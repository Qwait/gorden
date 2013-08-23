package gorden

var strategies = make(map[string]Strategy)

func AddStrategy(name string, strategy Strategy) {
    strategies[name] = strategy
}
