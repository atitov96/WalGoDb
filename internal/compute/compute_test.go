package compute

import (
	"atitov96/walgodb/internal/storage"
	"atitov96/walgodb/pkg/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	p := NewParser()

	tests := []struct {
		name       string
		expression string
		want       Command
		wantErr    bool
	}{
		{"Valid SET command", "SET key value", Command{Type: "SET", Args: []string{"key", "value"}}, false},
		{"Valid GET command", "GET key", Command{Type: "GET", Args: []string{"key"}}, false},
		{"Valid DEL command", "DEL key", Command{Type: "DEL", Args: []string{"key"}}, false},
		{"Empty expression", "", Command{}, true},
		{"Unknown command", "UNKNOWN key", Command{}, true},
		{"Invalid arguments", "SET key! value!", Command{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(tt.expression)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestComputeLayer_Execute(t *testing.T) {
	log, _ := logger.NewLogger("info", "stdout")
	parser := NewParser()
	engine := storage.NewInMemoryEngine()
	compute := NewComputeLayer(parser, engine, log)

	tests := []struct {
		name    string
		command string
		want    string
		wantErr bool
	}{
		{"SET key", "SET test_key test_value", "OK", false},
		{"GET existing key", "GET test_key", "test_value", false},
		{"GET non-existing key", "GET non_existing_key", "NOT FOUND", false},
		{"DEL existing key", "DEL test_key", "OK", false},
		{"DEL non-existing key", "DEL non_existing_key", "OK", false},
		{"Invalid command", "INVALID test_key", "", true},
		{"Set with wrong number of arguments", "SET test_key", "", true},
		{"Get with wrong number of arguments", "GET key1 key2", "", true},
		{"Del with wrong number of arguments", "DEL key1 key2", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compute.Execute(tt.command)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestInMemoryEngine(t *testing.T) {
	engine := storage.NewInMemoryEngine()

	t.Run("Set and Get", func(t *testing.T) {
		engine.Set("test_key", "test_value")
		val, ok := engine.Get("test_key")
		assert.True(t, ok)
		assert.Equal(t, "test_value", val)
	})

	t.Run("Get non-existing key", func(t *testing.T) {
		_, ok := engine.Get("non_existing_key")
		assert.False(t, ok)
	})

	t.Run("Delete", func(t *testing.T) {
		engine.Set("test_key", "test_value")
		engine.Delete("test_key")
		_, ok := engine.Get("test_key")
		assert.False(t, ok)
	})
}

func BenchmarkSet(b *testing.B) {
	log, _ := logger.NewLogger("info", "stdout")
	parser := NewParser()
	engine := storage.NewInMemoryEngine()
	compute := NewComputeLayer(parser, engine, log)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = compute.Execute("SET key value")
	}
}

func BenchmarkGet(b *testing.B) {
	log, _ := logger.NewLogger("info", "stdout")
	parser := NewParser()
	engine := storage.NewInMemoryEngine()
	compute := NewComputeLayer(parser, engine, log)

	_, _ = compute.Execute("SET key value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = compute.Execute("GET key")
	}
}

func BenchmarkDel(b *testing.B) {
	log, _ := logger.NewLogger("info", "stdout")
	parser := NewParser()
	engine := storage.NewInMemoryEngine()
	compute := NewComputeLayer(parser, engine, log)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = compute.Execute("SET key value")
		_, _ = compute.Execute("DEL key")
	}
}
