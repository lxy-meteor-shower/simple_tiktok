update_user_kitex_gen:
	kitex -I idl/ -module demo/tiktok -type protobuf idl/user.proto

update_video_kitex_gen:
	kitex -I idl/ -module demo/tiktok -type protobuf idl/video.proto

update_interact_kitex_gen:
	kitex -I idl/ -module demo/tiktok -type protobuf idl/interact.proto
