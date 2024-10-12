function init() {
    const image = $('.book-card .book-image');

    const { base, count } = image.data();

    setImages(0, undefined);

    const leftArrow = $('.book-card .arrow.left');
    const rightArrow = $('.book-card .arrow.right');

    $('.book-card .booklet-mode').click(function() {
        image.toggleClass('booklet');
        $(this).toggleClass('enabled');
        invalidate();
    });

    $('.book-card .arrow').click(function() {
        if ($(this).hasClass('disabled')) {
            return;
        }

        const { page } = image.data();

        let newPage = page;

        if ($(this).is(leftArrow)) {
            if (page == 0) {
                return;
            }
            
            newPage--;
            if (image.hasClass('booklet')) {
                newPage--;
            }
        } else if ($(this).is('.right')) {
            if (page == count - 1) {
                return;
            }

            newPage++;
            if (image.hasClass('booklet')) {
                newPage++;
            }
        }

        image.data({ page: newPage });

        invalidate();
    });

    function invalidate() {
        const { page } = image.data();

        switch (page) {
            case 0:
                leftArrow.addClass('disabled');
                rightArrow.removeClass('disabled');
                break;
            
            case count-1:
                leftArrow.removeClass('disabled');
                rightArrow.addClass('disabled');
                break;

            default:
                leftArrow.removeClass('disabled');
                rightArrow.removeClass('disabled');
                break;
        }

        const [ pageLeft, pageRight ] = calcNewPages(page);

        setImages(pageLeft, pageRight);
    }

    function calcNewPages(page) {
        const isBookletMode = image.hasClass('booklet');

        if (page == 0 || !isBookletMode || page == count-1) {
            return [page];
        }

        return [page, page+1];
    }

    function setImages(pageLeft, pageRight) {
        image.find('img.left').attr({ src: `${base}/${pageLeft+1}.jpg` });

        if (pageRight) {
            image.find('img.right').show().attr({ src: `${base}/${pageRight+1}.jpg` });
        } else {
            image.find('img.right').hide();
        }
    }
}
