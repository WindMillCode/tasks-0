import { StyleSheet,View } from 'react-native';
import { useTranslation } from 'react-i18next';
import { useWMLTheme } from '@/constants/Theme';
import {useWMLTranslation } from '@/constants/Global';

export default function WMLTemplate({

}) {
  const {isDark,theme} = useWMLTheme();
  const styles = createStyles({
    isDark,
    theme
  });
  const { t} = useWMLTranslation();

  return (
    <View style={styles.container}>

    </View>
  );
}


function createStyles(props) {
  let {isDark,theme} = props
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


