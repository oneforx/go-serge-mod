package compositions

import "github.com/oneforx/go-serge/ecs"

var (
	MOVABLE_COMPOSITION  ecs.Composition                    = []string{"oneforx:position"}
	TURNABLE_COMPOSITION ecs.Composition                    = []string{"oneforx:orientation"}
	PLAYER_COMPOSITION   ecs.Composition                    = []string{"oneforx:position", "oneforx:dimension", "oneforx:solid"}
	BULLET_COMPOSITION   ecs.Composition                    = []string{"oneforx:position", "oneforx:dimension", "oneforx:solid", "oneforx:speed", "oneforx:force", "oneforx:distance", "oneforx:orientation"}
	Compositions         map[ecs.Identifier]ecs.Composition = map[ecs.Identifier]ecs.Composition{
		{Namespace: "oneforx", Path: "movable"}: []string{"oneforx:position"},
	}
)
