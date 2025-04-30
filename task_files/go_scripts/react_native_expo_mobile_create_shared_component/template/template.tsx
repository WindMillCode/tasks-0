import { StyleSheet,View } from 'react-native';
import { useWMLTheme } from '@/constants/Theme';
import {useWMLTranslation } from '@/constants/MyGlobal';

export default function WMLTemplate(props) {
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
    label:{
      color:theme.text
    },
    container: {
      flex: 1,
      backgroundColor:theme.background
    },
    mainView: {
      padding: 20
    },
  });
}


