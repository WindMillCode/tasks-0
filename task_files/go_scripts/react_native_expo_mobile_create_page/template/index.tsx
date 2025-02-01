import { StyleSheet } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useTranslation } from 'react-i18next';
import { useWMLTheme } from '@/constants/Theme';

export default function WMLTemplatePage() {
  const {colorScheme,theme} = useWMLTheme();
  const styles = createStyles({
    colorScheme,
    theme
  });
  const { t } = useTranslation();

  return (
    <SafeAreaView style={styles.container}>

    </SafeAreaView>
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


