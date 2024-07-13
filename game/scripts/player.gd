extends CharacterBody2D


const SPEED = 500.0
const JUMP_VELOCITY = -800.0

# Get the gravity from the project settings to be synced with RigidBody nodes.
var gravity = ProjectSettings.get_setting("physics/2d/default_gravity")

@export var spell: PackedScene

func _physics_process(delta):
	# Add the gravity.
	if not is_on_floor():
		velocity.y += gravity * delta

		# Get mouse position in world space
	var mouse_pos = get_global_mouse_position()
	var mdirection = (mouse_pos - $WandRotation.global_position).normalized()
	
	# Rotate the wand to face the mouse
	$WandRotation.rotation = mdirection.angle()
	
	# Handle jump.
	if Input.is_action_just_pressed("jump") and is_on_floor():
		velocity.y = JUMP_VELOCITY
	
	if Input.is_action_just_pressed("Fire"):
		var sp = spell.instantiate()  # Correctly instantiate the spell
		sp.global_position = $WandRotation.get_node("SpellSpawn").global_position
		sp.rotation_degrees = $WandRotation.rotation_degrees
		get_parent().add_child(sp)
	# Get the input direction and handle the movement/deceleration.
	# As good practice, you should replace UI actions with custom gameplay actions.
	var direction = Input.get_axis("run_left", "run_right")
	if direction:
		velocity.x = direction * SPEED
	else:
		velocity.x = move_toward(velocity.x, 0, SPEED)

	move_and_slide()
