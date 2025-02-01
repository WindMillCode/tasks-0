import { StyleSheet,View } from 'react-native';
import { useTranslation } from 'react-i18next';
import { useWMLTheme } from '@/constants/Theme';

export default function WMLTemplate({

}) {
  const {colorScheme,theme} = useWMLTheme();
  const styles = createStyles({
    colorScheme,
    theme
  });
  const { t } = useTranslation();

  return (
    <View style={styles.container}>

    </View>
  );
}


function createStyles(props) {
  let {colorScheme,theme} = props
  return StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor:theme.background
    },
    mainView: {
      padding: 20
    },
  });
}


