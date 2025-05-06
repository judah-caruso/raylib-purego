package internal

var (
	//------------------------------------------------------------------------------------
	// Audio Loading and Playing Functions (Module: audio)
	//------------------------------------------------------------------------------------

	// Audio device management functions
	InitAudioDevice    = bind(&void, "InitAudioDevice")         // Initialize audio device and context
	CloseAudioDevice   = bind(&void, "CloseAudioDevice")        // Close the audio device and context
	IsAudioDeviceReady = bind(&cbool, "IsAudioDeviceReady")     // Check if audio device has been initialized successfully
	SetMasterVolume    = bind(&void, "SetMasterVolume", &float) // volume :: Set master volume (listener)
	GetMasterVolume    = bind(&float, "GetMasterVolume")        // Get master volume (listener)

	// Wave/Sound loading/unloading functions
	LoadWave           = bind(&tWave, "LoadWave", &cstr)                         // fileName :: Load wave data from file
	LoadWaveFromMemory = bind(&tWave, "LoadWaveFromMemory", &cstr, &ustr, &cint) // fileType, fileData, dataSize :: Load wave from memory buffer, fileType refers to extension: i.e. '.wav'
	IsWaveValid        = bind(&cbool, "IsWaveValid", &tWave)                     // wave :: Checks if wave data is valid (data loaded and parameters)
	LoadSound          = bind(&tSound, "LoadSound", &cstr)                       // fileName :: Load sound from file
	LoadSoundFromWave  = bind(&tSound, "LoadSoundFromWave", &tWave)              // wave :: Load sound from wave data
	LoadSoundAlias     = bind(&tSound, "LoadSoundAlias", &tSound)                // source :: Create a new sound that shares the same sample data as the source sound, does not own the sound data
	IsSoundValid       = bind(&cbool, "IsSoundValid", &tSound)                   // sound :: Checks if a sound is valid (data loaded and buffers initialized)
	UpdateSound        = bind(&void, "UpdateSound", &tSound, &ptr, &cint)        // sound, data, sampleCount :: Update sound buffer with new data
	UnloadWave         = bind(&void, "UnloadWave", &tWave)                       // wave :: Unload wave data
	UnloadSound        = bind(&void, "UnloadSound", &tSound)                     // sound :: Unload sound
	UnloadSoundAlias   = bind(&void, "UnloadSoundAlias", &tSound)                // alias :: Unload a sound alias (does not deallocate sample data)
	ExportWave         = bind(&cbool, "ExportWave", &tWave, &cstr)               // wave, fileName :: Export wave data to file, returns true on success
	ExportWaveAsCode   = bind(&cbool, "ExportWaveAsCode", &tWave, &cstr)         // wave, fileName :: Export wave sample data to code (.h), returns true on success

	// Wave/Sound management functions
	PlaySound         = bind(&void, "PlaySound", &tSound)                         // sound :: Play a sound
	StopSound         = bind(&void, "StopSound", &tSound)                         // sound :: Stop playing a sound
	PauseSound        = bind(&void, "PauseSound", &tSound)                        // sound :: Pause a sound
	ResumeSound       = bind(&void, "ResumeSound", &tSound)                       // sound ::  Resume a paused sound
	IsSoundPlaying    = bind(&cbool, "IsSoundPlaying", &tSound)                   // sound :: Check if a sound is currently playing
	SetSoundVolume    = bind(&void, "SetSoundVolume", &tSound, &float)            // sound, volume :: Set volume for a sound (1.0 is max level)
	SetSoundPitch     = bind(&void, "SetSoundPitch", &tSound, &float)             // sound, pitch :: Set pitch for a sound (1.0 is base level)
	SetSoundPan       = bind(&void, "SetSoundPan", &tSound, &float)               // sound, pan :: Set pan for a sound (0.5 is center)
	WaveCopy          = bind(&tWave, "WaveCopy", &tWave)                          // wave :: Copy a wave to a new wave
	WaveCrop          = bind(&void, "WaveCrop", &tWavePtr, &cint, &cint)          // wave, initFrame, finalFrame :: Crop a wave to defined frames range
	WaveFormat        = bind(&void, "WaveFormat", &tWavePtr, &cint, &cint, &cint) // wave, sampleRate, sampleSize, channels :: Convert wave data to desired format
	LoadWaveSamples   = bind(&floatPtr, "LoadWaveSamples", &tWave)                // wave :: Load samples data from wave as a 32bit float data array
	UnloadWaveSamples = bind(&void, "UnloadWaveSamples", &floatPtr)               // samples :: Unload samples data loaded with LoadWaveSamples()

	// Music management functions
	LoadMusicStream           = bind(&tMusic, "LoadMusicStream", &cstr)                         // fileName :: Load music stream from file
	LoadMusicStreamFromMemory = bind(&tMusic, "LoadMusicStreamFromMemory", &cstr, &ustr, &cint) // fileType, data, dataSize :: Load music stream from data
	IsMusicValid              = bind(&cbool, "IsMusicValid", &tMusic)                           // music :: Checks if a music stream is valid (context and buffers initialized)
	UnloadMusicStream         = bind(&void, "UnloadMusicStream", &tMusic)                       // music :: Unload music stream
	PlayMusicStream           = bind(&void, "PlayMusicStream", &tMusic)                         // music :: Start music playing
	IsMusicStreamPlaying      = bind(&cbool, "IsMusicStreamPlaying", &tMusic)                   // music :: Check if music is playing
	UpdateMusicStream         = bind(&void, "UpdateMusicStream", &tMusic)                       // music :: Updates buffers for music streaming
	StopMusicStream           = bind(&void, "StopMusicStream", &tMusic)                         // music :: Stop music playing
	PauseMusicStream          = bind(&void, "PauseMusicStream", &tMusic)                        // music :: Pause music playing
	ResumeMusicStream         = bind(&void, "ResumeMusicStream", &tMusic)                       // music :: Resume playing paused music
	SeekMusicStream           = bind(&void, "SeekMusicStream", &tMusic, &float)                 // music, position :: Seek music to a position (in seconds)
	SetMusicVolume            = bind(&void, "SetMusicVolume", &tMusic, &float)                  // music, volume :: Set volume for music (1.0 is max level)
	SetMusicPitch             = bind(&void, "SetMusicPitch", &tMusic, &float)                   // music, pitch :: Set pitch for a music (1.0 is base level)
	SetMusicPan               = bind(&void, "SetMusicPan", &tMusic, &float)                     // music, pan :: Set pan for a music (0.5 is center)
	GetMusicTimeLength        = bind(&float, "GetMusicTimeLength", &tMusic)                     // music :: Get music time length (in seconds)
	GetMusicTimePlayed        = bind(&float, "GetMusicTimePlayed", &tMusic)                     // music :: Get current music time played (in seconds)

	// AudioStream management functions
	LoadAudioStream                 = bind(&tAudioStream, "LoadAudioStream", &ucint, &ucint, &ucint) // sampleRate, sampleSize, channels :: Load audio stream (to stream raw audio pcm data)
	IsAudioStreamValid              = bind(&cbool, "IsAudioStreamValid", &tAudioStream)              // stream :: Checks if an audio stream is valid (buffers initialized)
	UnloadAudioStream               = bind(&void, "UnloadAudioStream", &tAudioStream)                // stream :: Unload audio stream and free memory
	UpdateAudioStream               = bind(&void, "UpdateAudioStream", &tAudioStream, &ptr, &cint)   // stream, data, frameCount :: Update audio stream buffers with data
	IsAudioStreamProcessed          = bind(&cbool, "IsAudioStreamProcessed", &tAudioStream)          // stream :: Check if any audio stream buffers requires refill
	PlayAudioStream                 = bind(&void, "PlayAudioStream", &tAudioStream)                  // stream :: Play audio stream
	PauseAudioStream                = bind(&void, "PauseAudioStream", &tAudioStream)                 // stream :: Pause audio stream
	ResumeAudioStream               = bind(&void, "ResumeAudioStream", &tAudioStream)                // stream :: Resume audio stream
	IsAudioStreamPlaying            = bind(&cbool, "IsAudioStreamPlaying", &tAudioStream)            // stream :: Check if audio stream is playing
	StopAudioStream                 = bind(&void, "StopAudioStream", &tAudioStream)                  // stream :: Stop audio stream
	SetAudioStreamVolume            = bind(&void, "SetAudioStreamVolume", &tAudioStream, &float)     // stream, volume :: Set volume for audio stream (1.0 is max level)
	SetAudioStreamPitch             = bind(&void, "SetAudioStreamPitch", &tAudioStream, &float)      // stream, pitch :: Set pitch for audio stream (1.0 is base level)
	SetAudioStreamPan               = bind(&void, "SetAudioStreamPan", &tAudioStream, &float)        // stream, pan :: Set pan for audio stream (0.5 is centered)
	SetAudioStreamBufferSizeDefault = bind(&void, "SetAudioStreamBufferSizeDefault", &cint)          // stream, size :: Default size for new audio streams

	// SetAudioStreamCallback     = bind(&void, "SetAudioStreamCallback", &tAudioStream, AudioCallback)     // stream, callback :: Audio thread callback to request new data
	// AttachAudioStreamProcessor = bind(&void, "AttachAudioStreamProcessor", &tAudioStream, AudioCallback) // stream, processor :: Attach audio stream processor to stream, receives the samples as 'float'
	// DetachAudioStreamProcessor = bind(&void, "DetachAudioStreamProcessor", &tAudioStream, AudioCallback) // stream, processor :: Detach audio stream processor from stream
	// AttachAudioMixedProcessor  = bind(&void, "AttachAudioMixedProcessor", AudioCallback)                 // processor :: Attach audio stream processor to the entire audio pipeline, receives the samples as 'float'
	// DetachAudioMixedProcessor  = bind(&void, "DetachAudioMixedProcessor", AudioCallback)                 // processor :: Detach audio stream processor from the entire audio pipeline
)
