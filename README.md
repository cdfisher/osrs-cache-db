# Simple OSRS Cache DB

Dumps cache definitions for items, NPCs, and objects to a SQLite3 DB and includes a simple server to query the DB.

## API
Build and run the package contained in `/server/`. By default this makes the API available on localhost:8080.

URLs follow the format `http://localhost:8080/<def>/<key>/<value>`.

For example, to get the cache definition for Fancy boots, I would go to `http://localhost:8080/items/name/Fancy boots`

### Defs & Keys

This currently supports cache definitions for items, npcs, and objects. Supported keys for each are as follows:

* `items`:
    * `id`

    * `name`

    * `examine`

    * `resize_x`

    * `resize_y`

    * `resize_z`

    * `xan2d`

    * `yan2d`

    * `zan2d`

    * `cost`

    * `is_tradable`

    * `stackable`

    * `inventory_model`

    * `wear_pos_1`

    * `wear_pos_2`

    * `wear_pos_3`

    * `members`

    * `zoom_2d`

    * `x_offset_2d`

    * `y_offset_2d`

    * `ambient`

    * `contrast`

    * `options`

    * `interface_options`

    * `male_model_0`

    * `male_model_1`

    * `male_model_2`

    * `male_offset`

    * `male_head_model`

    * `male_head_model_2`

    * `female_model_0`

    * `female_model_1`

    * `female_model_2`

    * `female_offset`

    * `female_head_model`

    * `female_head_model_2`

    * `noted_id`

    * `noted_template`

    * `team`

    * `weight`

    * `shift_click_drop_index`

    * `bought_id`

    * `bought_template_id`

    * `placeholder_id`

    * `placeholder_template_id`

    * `color_find`

    * `color_replace`

    * `params`

    * `count_co`

    * `count_obj`

    * `texture_find`

    * `texture_replace`

    * `category`


* `npcs`:
    * `id`

    * `name`

    * `size`

    * `models`

    * `chathead_models`

    * `standing_animation`,

    * `idle_rotate_left_animation`

    * `idle_rotate_right_animation`

    * `walking_animation`

    * `rotate_left_animation`

    * `rotate_right_animation`

    * `run_animation`

    * `run_rotate_180_animation`

    * `run_rotate_left_animation`

    * `run_rotate_right_animation`

    * `crawl_animation`

    * `crawl_rotate_180_animation`

    * `crawl_rotate_left_animation`

    * `crawl_rotate_right_animation`

    * `actions`

    * `is_minimap_visible`

    * `combat_level`

    * `width_scale`

    * `height_scale`

    * `has_render_priority`

    * `ambient`

    * `contrast`

    * `head_icon_sprite_index`

    * `head_icon_archive_ids`

    * `rotation_speed`

    * `varbit_id`

    * `varp_index`

    * `is_interactable`

    * `rotation_flag`

    * `is_pet`

    * `configs`

    * `params`

    * `category`

    * `recolor_to_find`

    * `recolor_to_replace`

    * `retexture_to_find`

    * `retexture_to_replace`

    * `is_follower`

    * `low_priority_follower_ops`


* `objects`:
    * `id`

    * `name`

    * `decor_displacement`

    * `is_hollow`

    * `object_models`

    * `object_types`

    * `map_area_id`

    * `size_x`

    * `size_y`

    * `offset_x`

    * `offset_y`

    * `offset_height`

    * `merge_normals`

    * `wall_or_door`

    * `animation_id`

    * `varbit_id`

    * `ambient`

    * `contrast`

    * `recolor_to_find`

    * `recolor_to_replace`

    * `retexture_to_find`

    * `texture_to_replace`

    * `actions`

    * `interact_type`

    * `map_scene_id`

    * `blocking_mask`

    * `shadow`

    * `model_size_x`

    * `model_size_y`

    * `model_size_height`

    * `object_id`

    * `obstructs_ground`

    * `contoured_ground`

    * `supports_items`

    * `config_change_dest`

    * `category`

    * `is_rotated`

    * `varp_id`

    * `ambient_sound_id`

    * `ambient_sound_ids`

    * `ambient_sound_retain`

    * `ambient_sound_distance`

    * `ambient_sound_change_ticks_min`

    * `ambient_sound_change_ticks_max`

    * `params`

    * `a_bool_2111`

    * `blocks_projectile`

    * `randomize_anim_start`