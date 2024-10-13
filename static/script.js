function init() {
    const image = $('.book-card .book-image');

    const leftArrow = $('.book-card .arrow.left');
    const rightArrow = $('.book-card .arrow.right');

    const { base, count } = image.data();

    setImages(0, undefined);

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
        setImageSrc(image.find('img.left'), `${base}/${pageLeft+1}.jpg`);

        if (pageRight) {
            image.find('img.right').show();
            setImageSrc(image.find('img.right'), `${base}/${pageRight+1}.jpg`);
        } else {
            image.find('img.right').hide();
        }
    }

    function setImageSrc(image, src) {
        setLoading(true);
        image.attr({ src });
    }

    function setLoading(loading) {
        if (loading) {
            $('.book-card').addClass('loading');
            leftArrow.addClass('disabled');
            rightArrow.addClass('disabled');
        } else {
            $('.book-card').removeClass('loading');
            leftArrow.removeClass('disabled');
            rightArrow.removeClass('disabled');
        }
    }

    function setLoadingError(image) {
        setImageSrc(image, '/static/error.png');
        setLoading(false);
    }

    image.find('img').
        on('load', function() {
            if (this.naturalHeight + this.naturalWidth == 0) {
                setLoadingError($(this));
                return;
            }

            if (!image.hasClass('booklet') || $(this).hasClass('right')) {
                setLoading(false);
            }
        }).
        on('error', function() {
            setLoadingError($(this));
        });
}
