function init() {
    const img = $('.book-card .book-image img');

    const { base, count } = img.data();

    img.attr({ src: `${base}/1.jpg` });

    const leftArrow = $('.book-card .arrow.left');
    const rightArrow = $('.book-card .arrow.right');

    $('.book-card .arrow').click(function() {
        if ($(this).hasClass('disabled')) {
            return;
        }

        const { page } = img.data();

        let newPage = page;

        if ($(this).is(leftArrow)) {
            if (page == 0) {
                return;
            }
            
            newPage--;

            rightArrow.removeClass('disabled');
            if (newPage == 0) {
                leftArrow.addClass('disabled');
            } else {
                leftArrow.removeClass('disabled');
            }
        } else if ($(this).is('.right')) {
            if (page == count - 1) {
                return;
            }

            newPage++;

            leftArrow.removeClass('disabled');
            if (newPage == count - 1) {
                rightArrow.addClass('disabled');
            } else {
                rightArrow.removeClass('disabled');
            }
        }

        img.data({ page: newPage });
        img.attr({ src: `${base}/${newPage+1}.jpg` });
    });
}
