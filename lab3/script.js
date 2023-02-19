let music = [
        { artist: 'Фираз Шатохин', title: 'Положение (Скриптонит кавер)', duration: '3:33', file: '1.mp3' },
        { artist: 'Klywoodjeb', title: 'Bebra', duration: '1:10', file: '2.mp3' },
        { artist: 'PLAYBOI CARTI', title: 'WOK', duration: '2:04', file: '3.mp3' },
         ];
      
      let currentTrackIndex = 0;
      let audioElement = document.getElementById('audio');
      let titleElement = document.getElementById('title');
      let artistElement = document.getElementById('artist');
      let playPauseButton = document.getElementById('play-pause');
      let prevButton = document.getElementById('prev');
      let nextButton = document.getElementById('next');
      
      function playTrack() {
        audioElement.setAttribute('src', music[currentTrackIndex].file);
        audioElement.play();
        titleElement.innerText = music[currentTrackIndex].title;
        artistElement.innerText = music[currentTrackIndex].artist;
        playPauseButton.innerText = 'pause';
      }
      
      function pauseTrack() {
        audioElement.pause();
        playPauseButton.innerText = 'play';
      }
      
      function playPauseTrack() {
        if (audioElement.paused) {
          playTrack();
        } else {
          pauseTrack();
        }
      }
      
      function playPrevTrack() {
        currentTrackIndex--;
        if (currentTrackIndex < 0) {
          currentTrackIndex = music.length - 1;
        }
        playTrack();
      }
      
      function playNextTrack() {
        currentTrackIndex++;
        if (currentTrackIndex >= music.length) {
          currentTrackIndex = 0;
        }
        playTrack();
      }
      
      playTrack();
      
      playPauseButton.addEventListener('click', playPauseTrack);
      prevButton.addEventListener('click', playPrevTrack);
      nextButton.addEventListener('click', playNextTrack);