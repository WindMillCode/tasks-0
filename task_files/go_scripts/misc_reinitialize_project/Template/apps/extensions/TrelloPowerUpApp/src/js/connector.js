var onBtnClick = function (t, opts) {
  return t.popup({
    title:"Main PowerUp Pane",
    url:"pane.html",
    height: 600,
  })
};


TrelloPowerUp.initialize({
  'board-buttons': function (t, options) {
    return [
      {
        // we can either provide a button that has a callback function
        icon: "https://[PROJECT_NAME].com/trello-powerup-nav-icon.svg",
        text: 'Sprints',
        callback: onBtnClick,
        condition: 'edit'
      },
    ];
  }
});


