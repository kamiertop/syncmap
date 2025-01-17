package syncmap

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkSyncMap_Int(b *testing.B) {
	var m sync.Map
	b.Run("SyncMap_Store_int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}
	})
	b.Run("SyncMap_Load_int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, ok := m.Load(i)
			if ok {
				_ = value.(int)
			}
		}
	})
}

func BenchmarkSyncMap_String(b *testing.B) {
	var m sync.Map
	b.Run("SyncStore_string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Store(fmt.Sprintf("key-%d", i), fmt.Sprintf("value-%d", i))
		}
	})
	b.Run("SyncMap_Load_string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v, ok := m.Load(fmt.Sprintf("key-%d", i))
			if ok {
				_ = v.(string)
			}
		}
	})
}

func BenchmarkSyncMap_Struct(b *testing.B) {
	type customStruct struct {
		ID   int
		Name string
	}
	var m sync.Map
	b.Run("SyncMap_Store_struct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Store(i, customStruct{ID: i, Name: fmt.Sprintf("Name-%d", i)})
		}
	})
	b.Run("SyncMap_Load_struct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Load(i)
		}
	})
}

func BenchmarkGenericMap_Int(b *testing.B) {
	var m Map[int, int]
	b.Run("Generic_Map_Store_int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}
	})
	b.Run("Generic_Map_Load_int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Load(i)
		}
	})
}

func BenchmarkGenericMap_String(b *testing.B) {
	var m Map[string, string]
	b.Run("Generic_Map_Store_string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Store(fmt.Sprintf("key-%d", i), fmt.Sprintf("value-%d", i))
		}
	})
	b.Run("Generic_Map_Load_string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Load(fmt.Sprintf("key-%d", i))
		}
	})
}

func BenchmarkGenericMap_Struct(b *testing.B) {
	type customStruct struct {
		ID   int
		Name string
	}
	var m Map[int, customStruct]
	b.Run("Generic_Map_Store_struct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Store(i, customStruct{ID: i, Name: fmt.Sprintf("Name-%d", i)})
		}
	})
	b.Run("Generic_Map_Load_struct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Load(i)
		}
	})
}
