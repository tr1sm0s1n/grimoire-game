extends CharacterBody2D


const SPEED = 500.0
const JUMP_VELOCITY = -800.0

# Get the gravity from the project settings to be synced with RigidBody nodes.
var gravity = ProjectSettings.get_setting("physics/2d/default_gravity")

@onready var animated_sprite = $AnimatedSprite2D2
@onready var camera: Camera2D = $Camera2D

@export var spell: PackedScene

var owner_id = 1

func _enter_tree():
	# Set the owner_id to the unique ID of the local player
	#owner_id = multiplayer.get_unique_id()
	owner_id = name.to_int()
	set_multiplayer_authority(owner_id)
	print("Owner ID: ", owner_id)

func _ready():
	# Activate the camera only for the local player
	if camera != null:
		if owner_id == multiplayer.get_unique_id():
			camera.make_current()
			print("Camera activated for player with ID: ", owner_id)
	else:
		print("Camera node not found")

func _physics_process(delta):
	if multiplayer.multiplayer_peer == null:
		return
		
	if owner_id != multiplayer.get_unique_id():
		return
		
	# Add the gravity.
	if not is_on_floor():
		velocity.y += gravity * delta

		# Get mouse position in world space
	var mouse_pos = get_global_mouse_position()
	var mdirection = (mouse_pos - $WandRotation.global_position).normalized()
	
	# Rotate the wand to face the mouse
	$WandRotation.rotation = mdirection.angle()
	
	var is_falling = velocity.y > 0.0 and not is_on_floor()
	var is_jumping = Input.is_action_just_pressed("jump") and is_on_floor()
	var is_idle = is_on_floor() and is_zero_approx(velocity.x)
	var is_running = is_on_floor() and not is_zero_approx(velocity.x)
	
	if is_jumping:
		animated_sprite.play("jump")
	elif is_running:
		animated_sprite.play("run")
	elif is_falling:
		animated_sprite.play("fall")
	elif is_idle:
		animated_sprite.play("idle")
		
	
	# Handle jump.
	if is_jumping:
		velocity.y = JUMP_VELOCITY
	
	if Input.is_action_just_pressed("Fire"):
		var sp = spell.instantiate()  # Correctly instantiate the spell
		sp.global_position = $WandRotation.get_node("SpellSpawn").global_position
		sp.rotation_degrees = $WandRotation.rotation_degrees
		get_parent().add_child(sp)
	# Get the input direction and handle the movement/deceleration.
	# As good practice, you should replace UI actions with custom gameplay actions.
	var direction = Input.get_axis("run_left", "run_right")
	
		# Flip the sprite
	if direction > 0:
		animated_sprite.flip_h = false
	elif direction < 0:
		animated_sprite.flip_h = true
		
	if direction:
		velocity.x = direction * SPEED
	else:
		velocity.x = move_toward(velocity.x, 0, SPEED)

	move_and_slide()
