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

/*
// Music management functions
RLAPI Music LoadMusicStream(const char *fileName);                    // Load music stream from file
RLAPI Music LoadMusicStreamFromMemory(const char *fileType, const unsigned char *data, int dataSize); // Load music stream from data
RLAPI bool IsMusicValid(Music music);                                 // Checks if a music stream is valid (context and buffers initialized)
RLAPI void UnloadMusicStream(Music music);                            // Unload music stream
RLAPI void PlayMusicStream(Music music);                              // Start music playing
RLAPI bool IsMusicStreamPlaying(Music music);                         // Check if music is playing
RLAPI void UpdateMusicStream(Music music);                            // Updates buffers for music streaming
RLAPI void StopMusicStream(Music music);                              // Stop music playing
RLAPI void PauseMusicStream(Music music);                             // Pause music playing
RLAPI void ResumeMusicStream(Music music);                            // Resume playing paused music
RLAPI void SeekMusicStream(Music music, float position);              // Seek music to a position (in seconds)
RLAPI void SetMusicVolume(Music music, float volume);                 // Set volume for music (1.0 is max level)
RLAPI void SetMusicPitch(Music music, float pitch);                   // Set pitch for a music (1.0 is base level)
RLAPI void SetMusicPan(Music music, float pan);                       // Set pan for a music (0.5 is center)
RLAPI float GetMusicTimeLength(Music music);                          // Get music time length (in seconds)
RLAPI float GetMusicTimePlayed(Music music);                          // Get current music time played (in seconds)

// AudioStream management functions
RLAPI AudioStream LoadAudioStream(unsigned int sampleRate, unsigned int sampleSize, unsigned int channels); // Load audio stream (to stream raw audio pcm data)
RLAPI bool IsAudioStreamValid(AudioStream stream);                    // Checks if an audio stream is valid (buffers initialized)
RLAPI void UnloadAudioStream(AudioStream stream);                     // Unload audio stream and free memory
RLAPI void UpdateAudioStream(AudioStream stream, const void *data, int frameCount); // Update audio stream buffers with data
RLAPI bool IsAudioStreamProcessed(AudioStream stream);                // Check if any audio stream buffers requires refill
RLAPI void PlayAudioStream(AudioStream stream);                       // Play audio stream
RLAPI void PauseAudioStream(AudioStream stream);                      // Pause audio stream
RLAPI void ResumeAudioStream(AudioStream stream);                     // Resume audio stream
RLAPI bool IsAudioStreamPlaying(AudioStream stream);                  // Check if audio stream is playing
RLAPI void StopAudioStream(AudioStream stream);                       // Stop audio stream
RLAPI void SetAudioStreamVolume(AudioStream stream, float volume);    // Set volume for audio stream (1.0 is max level)
RLAPI void SetAudioStreamPitch(AudioStream stream, float pitch);      // Set pitch for audio stream (1.0 is base level)
RLAPI void SetAudioStreamPan(AudioStream stream, float pan);          // Set pan for audio stream (0.5 is centered)
RLAPI void SetAudioStreamBufferSizeDefault(int size);                 // Default size for new audio streams
RLAPI void SetAudioStreamCallback(AudioStream stream, AudioCallback callback); // Audio thread callback to request new data

RLAPI void AttachAudioStreamProcessor(AudioStream stream, AudioCallback processor); // Attach audio stream processor to stream, receives the samples as 'float'
RLAPI void DetachAudioStreamProcessor(AudioStream stream, AudioCallback processor); // Detach audio stream processor from stream

RLAPI void AttachAudioMixedProcessor(AudioCallback processor); // Attach audio stream processor to the entire audio pipeline, receives the samples as 'float'
RLAPI void DetachAudioMixedProcessor(AudioCallback processor); // Detach audio stream processor from the entire audio pipeline
*/
)
