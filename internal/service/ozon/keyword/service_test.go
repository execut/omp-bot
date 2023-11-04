package keyword

import (
    "gotest.tools/v3/assert"
    "testing"
)

func TestList(t *testing.T) {
    service := NewService()
    t.Parallel()
    t.Run("List return 5 items", func(t *testing.T) {
        t.Parallel()

        list := service.List()

        assert.Equal(t, 5, len(list))
    })
}

func TestDelete(t *testing.T) {
    service := NewService()
    t.Parallel()
    t.Run("Delete deleted one item", func(t *testing.T) {
        t.Parallel()

        service.Delete(2)

        list := service.List()
        assert.Equal(t, 4, len(list))
    })
}
