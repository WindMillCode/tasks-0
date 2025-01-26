import { StyleSheet,Appearance,View } from 'react-native';
import { MyColors } from '@/constants/Colors';
import { useTranslation } from 'react-i18next';
import { captureException } from "@sentry/react-native";
import { useWMLNavigation } from "@/constants/Nav";

export default function WMLTemplate({

}) {
  const colorScheme = Appearance.getColorScheme()
  const theme = colorScheme === 'dark' ? MyColors.dark : MyColors.light
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

  return StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor:props.theme.background
    },
    mainView: {
      padding: 20
    },
  });
}


