type LanguageDisplayProps = {
    languages: { [key: string]: string };
    };

    function LanguageDisplay({ languages }: LanguageDisplayProps) {

        const languageArray = Object.entries(languages).map(
            ([code, name]) => `${name} (${code})`
        );

        const languageString = languageArray.join(', ');

        return <span>{languageString}</span>;
    }

    export default LanguageDisplay;