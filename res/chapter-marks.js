/*******************************************************************************
* Copyright 2021 Stefan Majewsky <majewsky@gmx.net>
* SPDX-License-Identifier: GPL-3.0
*******************************************************************************/

(function() {
  const $ = s => document.querySelector(s);
  const $$ = s => document.querySelectorAll(s);

  //This file only gets loaded in /noises/ for the pages of individual episodes.
  //It only deals with interactivity of the chapter mark section, so if we
  //don't have chapter marks, we're done here.
  if ($('#chapters') === null) {
    return;
  }

  const audio = $('audio');

  //The chapter marks section is initially hidden. We show it when the user
  //starts playing the audio file.
  audio.addEventListener('play', event => {
    $('#chapters').style.display = 'block';
  });
  if (!audio.paused) {
    //playback may have started before the script was fully loaded
    $('#chapters').style.display = 'block';
  }

  //While the audio is playing, continuously animate the progress bar in the
  //chapter mark section.
  let previousTime = 0;
  const nextFrame = () => {
    //only adjust progress bar states if playback is progressing
    const currSecs = audio.currentTime;
    if (previousTime == currSecs) {
      return;
    }
    previousTime = currSecs;

    //update display of each chapter
    for (const chapter of $$('div.chapter')) {
      const startSecs = parseFloat(chapter.dataset.start);
      const endSecs = parseFloat(chapter.dataset.end);

      const filledBar = chapter.querySelector('.chapter-progress-filler');
      if (currSecs < startSecs) {
        //we have not reached this chapter yet
        filledBar.style.width = '0%';
        chapter.classList.remove('current');
      } else if (currSecs >= endSecs) {
        //we have finished this chapter already
        filledBar.style.width = '100%';
        chapter.classList.remove('current');
      } else {
        //this is the current chapter
        const progress = 100 * (currSecs - startSecs) / (endSecs - startSecs);
        filledBar.style.width = progress + '%';
        chapter.classList.add('current');
      }
    }
  };

  //Call nextFrame in a loop.
  let loopFrame;
  loopFrame = () => {
    nextFrame();
    window.requestAnimationFrame(loopFrame);
  }
  window.requestAnimationFrame(loopFrame);

  //The span.chapter-start are styled as links. Clicking on them skips to the
  //start of that chapter.
  const skipToChapterStart = (event) => {
    const startSecs = parseInt(event.target.parentElement.dataset.start, 10);
    audio.currentTime = startSecs;
  };
  for (const span of $$('span.chapter-start')) {
    span.addEventListener('click', skipToChapterStart);
  }
})();
