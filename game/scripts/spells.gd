extends CharacterBody2D

const SPEED = 500.0

# Get the gravity from the project settings to be synced with RigidBody nodes.
var gravity = ProjectSettings.get_setting("physics/2d/default_gravity")
var direaction: Vector2

func _ready():
	direaction = Vector2(1,0).rotated(rotation)

func _physics_process(delta):
	# Add the gravity.
	velocity = SPEED * direaction
	
	if not is_on_floor():
		velocity.y += gravity * 1 * delta

	move_and_slide()
	
	var collision_info = move_and_collide(velocity * delta)
	
	# Check for collision
	if collision_info:
		queue_free() # Destroy the spell on collision
