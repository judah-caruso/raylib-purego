package raylib

import "github.com/judah-caruso/raylib-purego/internal"

var (
	//------------------------------------------------------------------------------------
	// Audio Loading and Playing Functions (Module: audio)
	//------------------------------------------------------------------------------------

	// Audio device management functions
	initAudioDevice    = internal.Bind(&void, "InitAudioDevice")         // Initialize audio device and context
	closeAudioDevice   = internal.Bind(&void, "CloseAudioDevice")        // Close the audio device and context
	isAudioDeviceReady = internal.Bind(&cbool, "IsAudioDeviceReady")     // Check if audio device has been initialized successfully
	setMasterVolume    = internal.Bind(&void, "SetMasterVolume", &float) // volume :: Set master volume (listener)
	getMasterVolume    = internal.Bind(&float, "GetMasterVolume")        // Get master volume (listener)

	// Wave/Sound loading/unloading functions
	loadWave           = internal.Bind(&tWave, "LoadWave", &cstr)                         // fileName :: Load wave data from file
	loadWaveFromMemory = internal.Bind(&tWave, "LoadWaveFromMemory", &cstr, &ustr, &cint) // fileType, fileData, dataSize :: Load wave from memory buffer, fileType refers to extension: i.e. '.wav'
	isWaveValid        = internal.Bind(&cbool, "IsWaveValid", &tWave)                     // wave :: Checks if wave data is valid (data loaded and parameters)
	loadSound          = internal.Bind(&tSound, "LoadSound", &cstr)                       // fileName :: Load sound from file
	loadSoundFromWave  = internal.Bind(&tSound, "LoadSoundFromWave", &tWave)              // wave :: Load sound from wave data
	loadSoundAlias     = internal.Bind(&tSound, "LoadSoundAlias", &tSound)                // source :: Create a new sound that shares the same sample data as the source sound, does not own the sound data
	isSoundValid       = internal.Bind(&cbool, "IsSoundValid", &tSound)                   // sound :: Checks if a sound is valid (data loaded and buffers initialized)
	updateSound        = internal.Bind(&void, "UpdateSound", &tSound, &ptr, &cint)        // sound, data, sampleCount :: Update sound buffer with new data
	unloadWave         = internal.Bind(&void, "UnloadWave", &tWave)                       // wave :: Unload wave data
	unloadSound        = internal.Bind(&void, "UnloadSound", &tSound)                     // sound :: Unload sound
	unloadSoundAlias   = internal.Bind(&void, "UnloadSoundAlias", &tSound)                // alias :: Unload a sound alias (does not deallocate sample data)
	exportWave         = internal.Bind(&cbool, "ExportWave", &tWave, &cstr)               // wave, fileName :: Export wave data to file, returns true on success
	exportWaveAsCode   = internal.Bind(&cbool, "ExportWaveAsCode", &tWave, &cstr)         // wave, fileName :: Export wave sample data to code (.h), returns true on success

	// Wave/Sound management functions
	playSound         = internal.Bind(&void, "PlaySound", &tSound)                         // sound :: Play a sound
	stopSound         = internal.Bind(&void, "StopSound", &tSound)                         // sound :: Stop playing a sound
	pauseSound        = internal.Bind(&void, "PauseSound", &tSound)                        // sound :: Pause a sound
	resumeSound       = internal.Bind(&void, "ResumeSound", &tSound)                       // sound ::  Resume a paused sound
	isSoundPlaying    = internal.Bind(&cbool, "IsSoundPlaying", &tSound)                   // sound :: Check if a sound is currently playing
	setSoundVolume    = internal.Bind(&void, "SetSoundVolume", &tSound, &float)            // sound, volume :: Set volume for a sound (1.0 is max level)
	setSoundPitch     = internal.Bind(&void, "SetSoundPitch", &tSound, &float)             // sound, pitch :: Set pitch for a sound (1.0 is base level)
	setSoundPan       = internal.Bind(&void, "SetSoundPan", &tSound, &float)               // sound, pan :: Set pan for a sound (0.5 is center)
	waveCopy          = internal.Bind(&tWave, "WaveCopy", &tWave)                          // wave :: Copy a wave to a new wave
	waveCrop          = internal.Bind(&void, "WaveCrop", &tWavePtr, &cint, &cint)          // wave, initFrame, finalFrame :: Crop a wave to defined frames range
	waveFormat        = internal.Bind(&void, "WaveFormat", &tWavePtr, &cint, &cint, &cint) // wave, sampleRate, sampleSize, channels :: Convert wave data to desired format
	loadWaveSamples   = internal.Bind(&floatPtr, "LoadWaveSamples", &tWave)                // wave :: Load samples data from wave as a 32bit float data array
	unloadWaveSamples = internal.Bind(&void, "UnloadWaveSamples", &floatPtr)               // samples :: Unload samples data loaded with LoadWaveSamples()

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
